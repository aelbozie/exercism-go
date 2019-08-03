class Matrix(object):
    def __init__(self, matrix_string):
        self.matrix_string = matrix_string.split('\n')

    def row(self, index):
        return list(map(int, self.matrix_string[index -1].split()))

    def column(self, index):
        column = []
        for i in range(len(self.matrix_string)):
            column.append(int(self.matrix_string[i].split()[index-1]))
        return column
