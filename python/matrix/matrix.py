class Matrix(object):
    def __init__(self, matrix_string):
        self.rows = [list(map(int, row.split(' ')))
                     for row in  matrix_string.split('\n')]
        
    def row(self, index):
        return self.rows[index - 1]

    def column(self, index):
        self.columns = [list(column) for column in zip(*self.rows)]
        return self.columns[index - 1]

