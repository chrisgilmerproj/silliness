#! /usr/local/bin/python

import math
import random


class Board(object):

    def __init__(self, size, gen_obj):
        self.size = size
        self.generator = self.board_generator(gen_obj)
        self.fill_board()

    def board_generator(self, val):
        count = 1
        output = val
        while count <= pow(self.size, 2):
            yield output
            output += val
            count += 1

    def fill_board(self):
        self.board = []
        for i in xrange(self.size):
            row = []
            for j in xrange(self.size):
                row.append(self.generator.next())
            self.board.append(row)

    def __getitem__(self, key):
        if not isinstance(key, int):
            raise TypeError
        return self.board[key]

    def __setitem__(self, key, val):
        if not isinstance(key, int):
            raise TypeError
        self.board[key] = val


class Player(object):

    def __init__(self, piece):
        self.piece = piece

    def print_move(self, move):
        print '\nAn %s was placed in position %d.' % (self.piece, move)

    def get_move(self, move_list):
        raise NotImplementedError


class HumanPlayer(Player):

    def get_move(self, board):
        square = pow(board.size, 2)
        while 1:
            try:
                move = int(raw_input('\nWhere to? '))
                if move in board.move_list:
                    print '\nPlease chose again.'
                elif move < 1 or move > square:
                    print '\nPlease enter a number between 1 and %d' % square
                else:
                    return move
            except ValueError:
                print '\nPlease input a valid number'


class RandomPlayer(Player):
    """
    This is an impementation of a random player
    """
    def get_move(self, board):
        return random.choice(board.get_open_moves())


class Game(object):

    def __init__(self, player1, player2, size=3):
        # Set Game Characteristics
        if size <= 1:
            raise Exception("Game size must be >= 2")
        else:
            self.size = int(size)
        self.square = pow(size, 2)
        self.move_list = []

        # Create the Boards
        self.board = Board(self.size, '')  # This stores the played pieces
        self.positions = Board(self.size, 1)  # This stores movement positions

        # Set the Players
        self.player1 = player1
        self.player2 = player2
        self.current_player = None

    def get_open_moves(self):
        return [x for x in xrange(1, self.square + 1) if x not in self.move_list]

    def print_intro(self):
        intro = 'Welcome to Tic-Tac-Toe.  Please make your move selection ' + \
                'by entering a number 1-%d corresponding to the movement ' + \
                'key on the right.'
        print intro % self.square

    def print_game(self):
        fill = int(math.ceil(math.log(self.square, 10)))
        print '\nBoard:' + '\t' * 3 + 'Movement Key:'
        for i in xrange(self.size):
            line = ''
            line += ' | '.join([str(j or ' ') for j in self.board[i]])
            line += '\t' * 2
            line += ' | '.join([str(j).zfill(fill) for j in self.positions[i]])
            print line

    def set_player(self):
        if not len(self.move_list) % 2:
            self.current_player = self.player1
        else:
            self.current_player = self.player2

    def get_move_position(self, move):
        move = move - 1
        row = move // self.size
        col = move % self.size
        return row, col

    def set_move(self, move):
        self.move_list.append(move)
        row, col = self.get_move_position(move)
        self.board[row][col] = self.current_player.piece

    def undo_last_move(self):
        move = self.move_list.pop()
        row, col = self.get_move_position(move)
        self.board[row][col] = ''

    def check_win(self):
        """
        Given the row and col of the piece we can reduce our search
        area to the row and column that was chosen or the diagonal if
        that applies.

        We can use a little math here to determine diagonal.
        One of two equations yields the answer:

            y = x + 0
            y = -x + self.size - 1

        """
        piece = self.current_player.piece
        row, col = self.get_move_position(self.move_list[-1])

        # Check diagonal
        win_diag = False
        if row - col == 0:
            win_diag = all([self.board[i][j] == piece \
                            for i in xrange(0, self.size) \
                            for j in xrange(0, self.size) \
                            if i == j])
        if row + col == self.size - 1:
            win_diag = all([self.board[i][j] == piece \
                            for i in xrange(0, self.size) \
                            for j in xrange(0, self.size) \
                            if i + j == self.size - 1])
        if win_diag:
            return True
        else:
            # look at row
            if all([self.board[row][i] == piece for i in xrange(0, self.size)]):
                return True
            # look at column
            if all([self.board[i][col] == piece for i in xrange(0, self.size)]):
                return True
        return False

    def check_draw(self):
        if len(self.move_list) == self.square:
            return True

    def play(self):
        self.print_intro()
        while 1:
            self.print_game()
            if self.move_list:
                if self.check_win():
                    print "\nPlayer %s has won the game" % self.current_player.piece
                    break
                elif self.check_draw():
                    print "\nThe game was a draw, no player wins"
                    break
            self.set_player()
            move = self.current_player.get_move(self)
            self.current_player.print_move(move)
            self.set_move(move)


def main():
    player1 = HumanPlayer('X')
    player2 = RandomPlayer('O')
    g = Game(player1, player2, size=3)
    g.play()


if __name__ == '__main__':
    try:
        main()
    except (KeyboardInterrupt, EOFError):
        print '\nSorry to see you go.  Game Ended'
