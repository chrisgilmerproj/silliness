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


class Game(object):

    def __init__(self, size=3):
        self.size = size
        self.square = pow(size, 2)
        self.move_list = []
        self.board = Board(size, '')
        self.positions = Board(size, 1)
        self.player1 = Player('X')
        self.player2 = Player('O')

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

    def print_move(self, move, player):
        person = 'You have'
        if not len(self.move_list) % 2:
            person = 'I will'
        print '\n%s put an %s in position %d.' % (person, player.piece, move)

    def set_move(self, move, player):
        move = move - 1
        self.board[move // self.size][move % self.size] = player.piece

    def check_win(self):
        pass

    def play(self):
        self.print_intro()
        while 1:
            self.print_game()
            player = self.get_player()
            move = player.get_move(self.size, self.move_list)
            self.print_move(move, player)
            self.set_move(move, player)
            self.check_win()


def main():
    g = Game()
    g.play()


if __name__ == '__main__':
    try:
        main()
    except (KeyboardInterrupt, EOFError):
        print '\nSorry to see you go.  Game Ended'
