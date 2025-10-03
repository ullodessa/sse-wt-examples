# Monitorsky Examples

This repo contains example client code for the [Monitorsky Realtime SaaS](https://monitorsky.cloud).

Our platform provides realtime messaging for developers with **Server-Sent Events (SSE)** today and **WebTransport** tomorrow.

## Examples

### JavaScript
```js
const es = new EventSource("https://api.sse-server.com/api/events?token=JWT_TOKEN");

es.onmessage = (event) => {
  console.log("Received:", JSON.parse(event.data));
};
