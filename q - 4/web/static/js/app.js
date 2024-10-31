// web/static/js/app.js

let ws;
let reconnectAttempts = 0;
const maxReconnectAttempts = 5;

function connectWebSocket() {
    ws = new WebSocket('ws://' + window.location.host + '/ws');
    
    ws.onopen = function() {
        console.log('Connected to server');
        document.getElementById('status').textContent = 'Connected';
        document.getElementById('status').className = 'success';
        reconnectAttempts = 0;
    };
    
    ws.onmessage = function(event) {
        const editor = document.getElementById('document-editor');
        editor.value = event.data;
    };
    
    ws.onclose = function() {
        document.getElementById('status').textContent = 'Disconnected';
        document.getElementById('status').className = 'error';
        
        if (reconnectAttempts < maxReconnectAttempts) {
            setTimeout(function() {
                reconnectAttempts++;
                connectWebSocket();
            }, 1000 * Math.pow(2, reconnectAttempts));
        }
    };
    
    ws.onerror = function(error) {
        console.error('WebSocket error:', error);
        ws.close();
    };
}

document.addEventListener('DOMContentLoaded', function() {
    connectWebSocket();
    
    const editor = document.getElementById('document-editor');
    let timeout = null;
    
    editor.addEventListener('input', function() {
        clearTimeout(timeout);
        timeout = setTimeout(function() {
            if (ws.readyState === WebSocket.OPEN) {
                ws.send(editor.value);
            }
        }, 100);
    });
});
