FROM nginx

ENV GREETER_SERVICE_ADDRESS=greeter-service:9090

COPY index.html /usr/share/nginx/html
COPY wasm_exec.js /usr/share/nginx/html
COPY main.wasm /usr/share/nginx/html