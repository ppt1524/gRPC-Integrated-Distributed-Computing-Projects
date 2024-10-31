# web/main.py

import sys
import os
import grpc
import asyncio
import argparse
import uuid
from fastapi import FastAPI, WebSocket, WebSocketDisconnect
from fastapi.responses import FileResponse
import uvicorn
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

app = FastAPI()

# Asynchronous gRPC channel
async def get_grpc_stub():
    channel = grpc.aio.insecure_channel('localhost:50053')
    await channel.channel_ready()
    stub = document_pb2_grpc.DocumentServiceStub(channel)
    logging.info('Connected to main server on port 50053')
    return stub

@app.get('/')
async def get():
    file_path = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'static', 'index.html')
    if not os.path.exists(file_path):
        logging.error(f'File not found: {file_path}')
        raise RuntimeError(f"File at path {file_path} does not exist.")
    return FileResponse(file_path)

@app.get('/static/{path:path}')
async def static_files(path):
    file_path = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'static', path)
    if not os.path.exists(file_path):
        logging.error(f'File not found: {file_path}')
        raise RuntimeError(f"File at path {file_path} does not exist.")
    return FileResponse(file_path)

@app.websocket('/ws')
async def websocket_endpoint(websocket: WebSocket):
    await websocket.accept()
    client_id = str(uuid.uuid4())
    logging.info(f"WebSocket connection accepted: {client_id}")

    stub = await get_grpc_stub()

    # Create a bidirectional stream
    stream = stub.StreamUpdates()

    async def send_updates():
        try:
            async for update in stream:
                logging.info(f'Received update from server: {update.content}')
                await websocket.send_text(update.content)
        except grpc.aio._call.AioRpcError as e:
            logging.error(f'gRPC StreamUpdates error: {e}')
        finally:
            await websocket.close()

    send_task = asyncio.create_task(send_updates())

    try:
        while True:
            data = await websocket.receive_text()
            logging.info(f'Received data from WebSocket: {data}')
            update_request = document_pb2.UpdateRequest(
                content=data,
                client_id=client_id
            )
            await stream.write(update_request)
            logging.info('Sent update to main server via gRPC')
    except WebSocketDisconnect:
        logging.info(f"WebSocket disconnected: {client_id}")
    except Exception as e:
        logging.error(f"WebSocket error: {e}")
    finally:
        await stream.done_writing()
        send_task.cancel()

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description="Web Interface for Live Document System")
    parser.add_argument('--port', type=int, default=8080, help='Port number to run the web interface on')
    args = parser.parse_args()
    uvicorn.run(app, host='0.0.0.0', port=args.port)
