from paseto.protocol.version4 import create_asymmetric_key, sign, verify

message = b"this is a public message"
public_key, secret_key = create_asymmetric_key()

token = sign(message, secret_key)
verified_message = verify(token, public_key)

assert verified_message == message
print(f"token={token}")
print(f"verified_message={verified_message}")
print(f"message={message}")