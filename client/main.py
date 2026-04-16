import json
import paho.mqtt.client as mqtt

BROKER = "localhost"
PORT   = 1883

switch = {
    "name": "salon-lewy",
    "state": False
}

def on_connect(client, userdata, flags, rc):
    if rc == 0:
        print(f"Connected to broker {BROKER}:{PORT}")
        client.subscribe("lightswitches/toggle")
        client.subscribe("lightswitches/new")
        print("Listening on: lightswitches/toggle, lightswitches/new")
    else:
        print(f"Connection failed, code: {rc}")

def on_message(client, userdata, msg):
    topic = msg.topic

    try:
        payload = json.loads(msg.payload.decode("utf-8"))
    except json.JSONDecodeError:
        print(f"Invalid JSON: {msg.payload}")
        return

    incoming_name = payload.get("name")

    if incoming_name is None:
        print("Missing 'name' field in payload")
        return

    if incoming_name != switch["name"]:
        print(f"Name mismatch: got '{incoming_name}', expected '{switch['name']}' - ignoring")
        return

    if topic == "lightswitches/toggle":
        new_state = payload.get("state", not switch["state"])
        switch["state"] = new_state
        print(f"Toggle -> state: {'ON' if switch['state'] else 'OFF'}")

    elif topic == "lightswitches/new":
        confirm = json.dumps({"name": switch["name"]})
        client.publish("lightswitches/new/confirm", confirm)
        print(f"Name matched -> sent confirm: {confirm}")

client = mqtt.Client(client_id="lightswitch-sim")
client.on_connect = on_connect
client.on_message = on_message

client.connect(BROKER, PORT, keepalive=60)
client.loop_forever()
