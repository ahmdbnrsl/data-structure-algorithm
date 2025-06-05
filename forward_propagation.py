import math

def sigmoid(x):
    return 1 / (1 + math.exp(-x))

input_vector = [1.0, 0.5, -1.5]

weights_input_hidden = [
    [0.2, 0.4, -0.5],
    [-0.3, 0.1, 0.2]
]
bias_hidden = [0.1, -0.2]

hidden_layer_output = []
for i in range(2):
    z = sum(input_vector[j] * weights_input_hidden[i][j] for j in range(3)) + bias_hidden[i]
    a = sigmoid(z)
    hidden_layer_output.append(a)

weights_hidden_output = [0.6, -0.1]
bias_output = 0.05

output_z = sum(hidden_layer_output[i] * weights_hidden_output[i] for i in range(2)) + bias_output
output = sigmoid(output_z)

print("Output:", round(output, 4))