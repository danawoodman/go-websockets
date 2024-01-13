let ws

function connect() {
  ws = new WebSocket("ws://localhost:9090/ws")
  ws.onopen = () => {
    console.log("connected")
    document.querySelector("#disconnected-msg").style.display = "none"
  }
  ws.onmessage = (e) => {
    const data = JSON.parse(e.data)
    console.log("received message:", data)
    document.querySelector("pre").textContent += e.data + "\n"
  }
  ws.onclose = () => {
    console.log("disconnected")
    document.querySelector("#disconnected-msg").style.display = "block"
    setTimeout(() => connect(), 1000)
  }
}

connect()

document.forms[0].onsubmit = (e) => {
  e.preventDefault()
  const input = document.querySelector("input")
  if (input.value) {
    console.log("sending message:", input.value)
    ws.send(JSON.stringify({ type: "message", data: input.value }))
    input.value = ""
  }
}
