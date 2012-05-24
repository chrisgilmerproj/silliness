#! /usr/local/bin/python

import math


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

    def get_move(self, size, move_list):
        raise NotImplementedError


class HumanPlayer(Player):

    def get_move(self, size, move_list):
        square = pow(size, 2)
        while 1:
            try:
                move = int(raw_input('\nWhere to? '))
                if move in move_list:
                    print '\nPlease chose again.'
                elif move < 1 or move > square:
                    print '\nPlease enter a number between 1 and %d' % square
                else:
                    move_list.append(move)
                    return move
            except ValueError:
                print '\nPlease input a valid number'


class ComputerPlayer(Player):

    def get_move(self, size, move_list):
        pass


class Game(object):

    def __init__(self, size=3):
        self.size = size
        self.square = pow(size, 2)
        self.move_list = []
        self.board = Board(size, '')
        self.positions = Board(size, 1)
        self.player1 = HumanPlayer('X')
        self.player2 = HumanPlayer('O')

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

    def get_player(self):
        if not len(self.move_list) % 2:
            return self.player1
        else:
            return self.player2

    def set_move(self, move, player):
        move = move - 1
        row = move // self.size
        col = move % self.size
        self.board[row][col] = player.piece
        return row, col

    def check_win(self, row, col, piece):
        """
        We can use a little math here to determine diagonal.
        One of two equations yields the answer:

            y = x + 0
            y = -x + self.size - 1

        """
        if row - col == 0:
            return all([self.board[i][j] == piece \
                        for i in xrange(0, self.size) \
                        for j in xrange(0, self.size) \
                        if i == j])
        elif row + col == self.size - 1:
            return all([self.board[i][j] == piece \
                        for i in xrange(0, self.size) \
                        for j in xrange(0, self.size) \
                        if i + j == self.size - 1])

    def check_draw(self):
        if len(self.move_list) == self.square:
            return True

    def play(self):
        player = None
        row, col = None, None
        self.print_intro()
        while 1:
            self.print_game()
            if self.move_list and player:
                if self.check_win(row, col, player.piece):
                    print "\nPlayer %s has won the game" % player.piece
                    break
                elif self.check_draw():
                    print "\nThe game was a draw, no player wins"
                    break
            player = self.get_player()
            move = player.get_move(self.size, self.move_list)
            player.print_move(move)
            row, col = self.set_move(move, player)


def main():
    g = Game()
    g.play()


if __name__ == '__main__':
    try:
        main()
    except (KeyboardInterrupt, EOFError):
        print '\nSorry to see you go.  Game Ended'
