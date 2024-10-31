# server/main.py

import sys
import os
import grpc
import asyncio
from concurrent import futures
import time
import logging

# Adjust the Python path to include the 'proto' directory
current_dir = os.path.dirname(os.path.abspath(__file__))
proto_dir = os.path.abspath(os.path.join(current_dir, '..', 'proto'))
if proto_dir not in sys.path:
    sys.path.append(proto_dir)

import document_pb2
import document_pb2_grpc

# Configure logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

class DocumentServer(document_pb2_grpc.DocumentServiceServicer):
    def __init__(self, logging_stub):
        self.content = ''
        self.clients = {}  # client_id -> asyncio.Queue
        self.client_counter = 0
        self.logging_stub = logging_stub
        self.lock = asyncio.Lock()

    async def StreamUpdates(self, request_iterator, context):
        async with self.lock:
            client_id = f'client-{self.client_counter}'
            self.client_counter += 1
            client_queue = asyncio.Queue()
            self.clients[client_id] = client_queue
            current_content = self.content
            logging.info(f'Client {client_id} connected')

        # Send current content to new client
        await context.write(document_pb2.UpdateResponse(
            content=current_content,
            client_id=client_id,
            success=True
        ))

        async def receive_updates():
            try:
                async for update in request_iterator:
                    async with self.lock:
                        self.content = update.content
                        logging.info(f'Received update from {client_id}: {update.content}')
                        # Broadcast to other clients
                        for cid, cq in self.clients.items():
                            if cid != client_id:
                                await cq.put(document_pb2.UpdateResponse(
                                    content=update.content,
                                    client_id=client_id,
                                    success=True
                                ))
                    # Log the update via Logging Service
                    await self.log_update(update)
            except Exception as e:
                logging.error(f'Client {client_id} disconnected: {e}')
            finally:
                async with self.lock:
                    if client_id in self.clients:
                        del self.clients[client_id]
                        logging.info(f'Client {client_id} removed')

        asyncio.create_task(receive_updates())

        # Send updates to this client
        try:
            while True:
                update_response = await client_queue.get()
                await context.write(update_response)
        except Exception as e:
            logging.error(f'Error sending updates to client {client_id}: {e}')
        finally:
            async with self.lock:
                if client_id in self.clients:
                    del self.clients[client_id]
                    logging.info(f'Client {client_id} removed')

    async def GetDocument(self, request, context):
        async with self.lock:
            content = self.content
        logging.info('GetDocument called')
        return document_pb2.GetDocumentResponse(content=content)

    async def log_update(self, update):
        try:
            response = await self.logging_stub.LogUpdates(update)
            logging.info(f'Logging response: {response.message} at {response.timestamp}')
        except grpc.aio._call.AioRpcError as e:
            logging.error(f'gRPC error while logging update: {e.code()} - {e.details()}')
        except Exception as e:
            logging.error(f'Unexpected error while logging update: {e}')

async def serve():
    # Connect to Logging Service
    channel = grpc.aio.insecure_channel('127.0.0.1:50052')  # Use IPv4
    await channel.channel_ready()
    logging_stub = document_pb2_grpc.DocumentServiceStub(channel)
    logging.info('Connected to logging service on port 50052')

    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))
    document_pb2_grpc.add_DocumentServiceServicer_to_server(
        DocumentServer(logging_stub), server)
    server.add_insecure_port('[::]:50053')
    await server.start()
    logging.info('Main server listening on port 50053')
    await server.wait_for_termination()

if __name__ == '__main__':
    try:
        asyncio.run(serve())
    except KeyboardInterrupt:
        logging.info("Main server stopped by user")
    except Exception as e:
        logging.error(f"Main server encountered an error: {e}")
