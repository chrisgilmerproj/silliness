#! /usr/local/bin/python

import itertools
import random


class Card(object):

    def __init__(self, rank, suit):
        self.value, self.rank = rank  # Tuple, value and representation
        self.suit = suit

    def __repr__(self):
        return "Card <%s, %s>" % (self.rank, self.suit)

    def __eq__(self, other):
        return self.value == other.value

    def __lt__(self, other):
        return self.value < other.value

    def __gt__(self, other):
        return self.value > other.value


class Deck(object):

    def __init__(self, rank_dict, suit_list):
        self.rank_dict = rank_dict
        self.suit_list = suit_list
        self.cards = []
        self.hands = []

    def __len__(self):
        return len(self.cards)

    def populate(self):
        rank_list = [(num, repr) for num, repr in self.rank_dict.iteritems()]
        self.cards = [Card(*card) for card in itertools.product(rank_list, self.suit_list)]

    def shuffle(self):
        random.shuffle(self.cards)

    def deal(self, num_players, deal_even=True):
        if deal_even and len(self.cards) % num_players != 0:
            print deal_even, len(self.cards) % num_players != 0
            raise Exception('With %d players you will not deal evenly' % num_players)
        hands = []
        for i, card in enumerate(self.cards):
            hand_len = i % num_players
            if len(hands) < hand_len + 1:
                hands.append([])
            hands[hand_len].append(card)
        self.hands = hands
        return self.hands


SUIT_LIST = ('spade', 'diamond', 'club', 'heart')
RANK_DICT = {
        2: '2',
        3: '3',
        4: '4',
        5: '5',
        6: '6',
        7: '7',
        8: '8',
        9: '9',
        10: '10',
        11: 'J',
        12: 'Q',
        13: 'K',
        14: 'A',
        }


def main():
    print "Let's play war!"

    # Create the deck
    deck = Deck(RANK_DICT, SUIT_LIST)
    deck.populate()
    deck.shuffle()
    deck.deal(2, deal_even=False)

    # Play the game!
    num_turns = 0
    while(all([len(hand) for hand in deck.hands])):
        winner, pile = compare_cards(*deck.hands)
        random.shuffle(list(set(pile)))

        deck.hands[winner].extend(pile)

        num_turns += 1
        print [len(hand) for hand in deck.hands]
        if num_turns > 10000:
            print "Too many turns, quitting game"
            break

    # Print winner and stats
    for i, hand in enumerate(deck.hands):
        if len(hand):
            winner = i + 1
    print "Congrats player %d" % winner
    print "Number of turns: %d" % num_turns


def compare_cards(deck1, deck2):
    c1 = deck1.pop(0)
    c2 = deck2.pop(0)

    winner = 0
    pile = [c1, c2]
    if c1 > c2:
        winner = 0
    if c1 < c2:
        winner = 1
    if c1 == c2:
        print '\tWar!'
        for deck in (deck1, deck2):
            if len(deck):
                pile.extend(war(deck))
        if len(deck1) == 0:
            deck1.append(c1)
        if len(deck2) == 0:
            deck2.append(c2)
        winner, pile2 = compare_cards(deck1, deck2)
        pile.extend(pile2)
    return winner, pile


def war(deck):
    pile = []
    for i in range(0, 3):
        if len(deck) > 1:
            pile.append(deck.pop(0))
    return pile


if __name__ == "__main__":
    main()
