import pandas as pd
import numpy as np
import matplotlib.pyplot as plt


def set_index(np: int, nc: int) -> int:
    """Parametrizes the np and nc into indexes

    Args:
        np (int): number of threads of producer
        nc (int): number of threads of consumer

    Returns:
        int: the parametrized index
    """
    result = {(1, 1): 0, (1, 2): 1, (1, 4): 2, (1, 8): 3, (1, 16): 4, (2, 1): 5, (4, 1): 6, (8, 1): 7, (16, 1): 8}

    return result[(np, nc)]


def parse_csv(data_file: str) -> pd.DataFrame:
    """Access the generated csv by the bash and treats it.

    Args:
        data_file (str): the path to the csv.

    Returns:
        pd.DataFrame: a pandas dataframe containing the treated data
    """
    df = pd.read_csv(data_file)
    df = df.groupby(['n', 'np', 'nc']).mean().reset_index()

    df['graph_index'] = df.apply(lambda x: set_index(
        np=x['np'], nc=x['nc']), axis=1)

    return df


def save_graph(image: pd.DataFrame, name: str) -> None:
    """ Receives a figure as dataframe and saves it as a png file

    Args:
        image (pd.DataFrame): the figure dataframe
        name (str): the figure name
    """
    image.get_figure().savefig(f'{name}.png')


def build_graph(df: pd.DataFrame) -> pd.DataFrame:
    """Receives a pandas dataframe and builds a graph of time x Number of threads parametrized

    Args:
        df (pd.DataFrame): the treated dataframe

    Returns:
        pd.DataFrame: the graph as dataframe
    """

    fig = df.pivot(index='graph_index', columns='n',
                   values='time').plot(title="Gráfico de Tempo x Número de Threads Produtor/Consumidor")
    fig.set_xlabel("Thread Produtor/Consumidor parametrizadas")
    fig.set_ylabel("Tempo (s)")
    return fig


if __name__ == "__main__":

    data_file = 'go_results.csv'

    df = parse_csv(data_file)

    image = build_graph(df)

    save_graph(image, "result")
