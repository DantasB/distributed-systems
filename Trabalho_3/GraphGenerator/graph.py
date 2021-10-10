import os
import datetime
import matplotlib.pyplot as plt

DATA_FOLDER = '../Results/'

def subtract_datetimes(first_line:str, last_line:str) -> int:
    """Receives 2 datetimes lines and returns the difference of it in seconds

    Args:
        first_line (str): the first line of the archive
        last_line (str): the last line of the archive

    Returns:
        int: number of seconds of the difference of dates
    """
    start_datetime = first_line.split(' Process')[0]
    end_datetime = last_line.split(' Process')[0]
    layout = '%Y/%m/%d %H:%M:%S.%f'
    start_datetime = datetime.datetime.strptime(start_datetime, layout)
    end_datetime = datetime.datetime.strptime(end_datetime, layout)
    return (end_datetime - start_datetime).total_seconds()

def get_txt_runtime(path:str) -> tuple:
    """Gets the number of processes and calculate the time that it took to run

    Args:
        path (str): the path of the run log

    Returns:
        tuple: a tuple containing the number of process and the runtime
    """
    txt_file = path.split('_')[1]
    processes_number = int(''.join(filter(str.isdigit, txt_file)))
    first_line = ''
    last_line = ''
    with open(path, encoding='utf-8') as file:
        lines = file.readlines()
        first_line = lines[0]
        last_line = lines[-1]
    runtime = subtract_datetimes(first_line, last_line)
    return processes_number, runtime

def get_folder_archives(data_folder:str) -> dict:
    """Iterates over the files in the results folder
    and returns a dictionary containing the tests cases and their outputs

    Args:
        data_folder (str): the path to the results folder

    Returns:
        dict: a dictionary where the key is the test case and the value is a list of the tests
    """
    output = {}
    test_case = 1
    for subdir, _, files in os.walk(data_folder):
        test_case_files = []
        if 'Case0' in subdir or subdir == '../Results/':
            continue
        for file in files:
            test_case_files.append(os.path.join(subdir, file))
        output[test_case] = test_case_files
        test_case += 1
    return output

def parse_texts(paths:list) -> tuple:
    """Gets the case_test name and calculates the runtime for each run in the case test

    Args:
        paths (list): a list containing all runs for a case_test

    Returns:
        tuple: a tuple containing the case_test number and the list of runtimes for each case test
    """
    result = []
    case_test = paths[0].split(os.sep)[2]
    for path in paths:
        result.append(get_txt_runtime(path))
    return case_test, result

def generate_graph(tests:dict) -> None:
    """Parses the tests dictionary and draw the referent graph

    Args:
        tests (dict): a dictionary containing all the paths for a test case
    """
    for case in tests:
        results = parse_texts(tests[case])
        draw_graph(results)

def save_graph(figure: plt.Figure, test_case:str) -> None:
    """Sets a subtitle and label names for a figure

    Args:
        figure (plt.Figure): the generated image
        test_case (str): the test case name
    """
    figure.suptitle(f'{test_case}: Gráfico de N x Tempo(s)')
    plt.xlabel('Número de Processos')
    plt.ylabel('Tempo(s)')
    plt.savefig(f'results/{test_case}.png')

def draw_graph(results:tuple) -> None:
    """Draws a graph for a given test_case

    Args:
        results (tuple): a tuple containing the test_case number and their results
    """
    test_case, results = results
    results.sort(key=lambda tup: tup[0])
    figure = plt.figure()
    plt.plot(*zip(*results))
    save_graph(figure, test_case)
    plt.clf()

if __name__ == "__main__":

    data = get_folder_archives(DATA_FOLDER)
    generate_graph(data)
