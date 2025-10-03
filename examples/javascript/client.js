// examples/javascript/client.js
const es = new EventSource("https://api.sse-server.com/api/events?token=JWT_TOKEN");

es.onmessage = (event) => {
  console.log("Received:", JSON.parse(event.data));
};
