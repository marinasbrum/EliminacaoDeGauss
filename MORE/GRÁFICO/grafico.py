import matplotlib.pyplot as plt
import pandas as pd

#  TESTES COM SEED ALEATÓRIA E TAMANHOS = 3, 6 E 10

data = {
    'N': [3, 6, 10],
    'Elapsed Time Go': [2.690, 4.828, 1.772],
    'CPU Time Go': [3.230, 5.363, 1.772],
    'Elapsed Time C': [0.126, 0.113, 0.362],
    'CPU Time C': [0.000, 0.000, 0.000],
    'Elapsed Time Rust': [0.123, 0.214, 0.293],
    'CPU Time Rust': [0.122, 0.214, 0.292]
}
df = pd.DataFrame(data)

plt.figure(figsize=(10, 5))
plt.subplot(1, 2, 1)
plt.plot(df['N'], df['Elapsed Time Go'], label='Go', marker='o')
plt.plot(df['N'], df['Elapsed Time C'], label='C', marker='o')
plt.plot(df['N'], df['Elapsed Time Rust'], label='Rust', marker='o')
plt.xlabel('Matrix Dimension N')
plt.ylabel('Elapsed Time (ms)')
plt.title('Elapsed Time Comparison')
plt.legend()

plt.subplot(1, 2, 2)
plt.plot(df['N'], df['CPU Time Go'], label='Go', marker='o')
plt.plot(df['N'], df['CPU Time C'], label='C', marker='o')
plt.plot(df['N'], df['CPU Time Rust'], label='Rust', marker='o')
plt.xlabel('Matrix Dimension N')
plt.ylabel('CPU Time (ms)')
plt.title('CPU Time Comparison')
plt.legend()
plt.tight_layout()
plt.show()

#  TESTES COM SEED = 12345 E TAMANHO = 5

data = {
    'Linguagem': ['Go', 'C', 'Rust'],
    'Elapsed Time': [3.163, 0.128, 0.110],
    'CPU Time': [3.688, 0.128, 0.000]
}
df = pd.DataFrame(data)

plt.figure(figsize=(12, 6))
indices = range(len(df))
bar_width = 0.50
plt.subplot(1, 2, 1)
plt.bar([i - bar_width/2 for i in indices], df['Elapsed Time'], bar_width, label='Elapsed Time', color='blue')
plt.xlabel('Linguagens', fontsize=12)
plt.ylabel('Elapsed Time (ms)', fontsize=12)
plt.title('Elapsed Time Comparison', fontsize=14)
plt.xticks(indices, df['Linguagem']) 
plt.legend()

plt.subplot(1, 2, 2)
plt.bar([i + bar_width/2 for i in indices], df['CPU Time'], bar_width, label='CPU Time', color='green')
plt.xlabel('Linguagens', fontsize=12)
plt.ylabel('CPU Time (ms)', fontsize=12)
plt.title('CPU Time Comparison', fontsize=14)
plt.xticks(indices, df['Linguagem'])  
plt.legend()
plt.tight_layout()
plt.show()
