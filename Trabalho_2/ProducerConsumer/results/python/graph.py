import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

DATA_FILE = 'go_results.csv'

df = pd.read_csv(DATA_FILE)
df = df.groupby(['n', 'np', 'nc']).mean().reset_index()
df.to_csv(DATA_FILE, sep=',', encoding='utf-8',
          header=True, decimal='.', float_format='%.4f', index=False)
tmp_df = df
n_sizes = df['n'].unique()
