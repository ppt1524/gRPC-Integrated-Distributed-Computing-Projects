# logging/main.py

import sys
import os
import grpc
from concurrent import futures
import asyncio
import time
import aiofiles
import logging

# Adjust the Python path to include the 'proto' directory
current_dir = os.path.dirname(os.path.abspath(__file__))
proto_dir = os.path.abspath(os.path.join(current_dir, '..', 'proto'))
if proto_dir not in sys.path:
    sys.path.append(proto_dir)

import document_pb2
import document_pb2_grpc

# Configure logging for the service itself
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

class LoggingService(document_pb2_grpc.DocumentServiceServicer):
    def __init__(self, log_file_path):
        self.log_file_path = log_file_path
        # Ensure the log file exists
        if not os.path.exists(self.log_file_path):
            open(self.log_file_path, 'w').close()

    async def LogUpdates(self, request, context):
        timestamp = time.strftime('%Y-%m-%dT%H:%M:%S%z', time.localtime())
        log_entry = f'[{timestamp}] Client {request.client_id} updated document: {request.content}\n'
        
        # Debug: Log received data
        logging.debug(f'Received LogUpdate - Client ID: "{request.client_id}", Content: "{request.content}"')
        
        if not request.client_id:
            logging.warning(f'LogUpdate received without client_id at {timestamp}')
        
        if not request.content:
            logging.warning(f'LogUpdate received without content from client {request.client_id} at {timestamp}')
        
        try:
            async with aiofiles.open(self.log_file_path, 'a') as log_file:
                await log_file.write(log_entry)
            logging.info(f'Logged update from {request.client_id}')
            response = document_pb2.LogResponse(
                message='Update logged successfully',
                timestamp=timestamp
            )
            return response
        except Exception as e:
            logging.error(f'Failed to write to log file: {e}')
            await context.abort(grpc.StatusCode.INTERNAL, f'Failed to write to log file: {e}')

async def serve():
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))
    document_pb2_grpc.add_DocumentServiceServicer_to_server(
        LoggingService('document_changes.log'), server)
    server.add_insecure_port('[::]:50052')
    await server.start()
    logging.info('Logging server listening on port 50052')
    await server.wait_for_termination()

if __name__ == '__main__':
    try:
        asyncio.run(serve())
    except KeyboardInterrupt:
        logging.info("Logging service stopped by user")
    except Exception as e:
        logging.error(f"Logging service encountered an error: {e}")
