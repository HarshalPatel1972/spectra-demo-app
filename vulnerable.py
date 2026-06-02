from cryptography.hazmat.primitives.asymmetric import ec
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes

# Highly vulnerable to Shor's Algorithm
private_key = ec.generate_private_key(ec.SECP256R1())

# Not vulnerable to Shor's, but key size halved by Grover's
key = b'32_byte_key_here...'
cipher = Cipher(algorithms.AES(key), modes.GCM(b'iv_here'))

print("Demo app initialized.")
