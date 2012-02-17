#! /usr/local/bin/python

import itertools
import random

# Constants used for the game
MAX_TURNS = 5000
PRINT_HANDS = True
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


class Card(object):
    """ A Card represents a suit, rank and value """

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
    """ A Deck represents a set of cards that can be played in a game """

    def __init__(self, rank_dict, suit_list):
        self.rank_dict = rank_dict
        self.suit_list = suit_list
        self.cards = []
        self.hands = []

    def __len__(self):
        """ Return the size of the deck """
        return len(self.cards)

    def populate(self):
        """ Populate the deck using the rank and suit information """
        rank_list = [(num, repr) for num, repr in self.rank_dict.iteritems()]
        self.cards = [Card(*card) for card in itertools.product(rank_list, self.suit_list)]

    def shuffle(self):
        """ Shuffle the cards in the deck """
        random.shuffle(self.cards)

    def deal(self, num_players, deal_even=True):
        """
        Deal the cards into several hands based upon number of players.

        @param num_players: The number of players to deal to
        @param deal_even: When set to false players will be delt all cards
            but the size of their hands may differ
        """
        if deal_even and len(self.cards) % num_players != 0:
            raise Exception('With %d players you will not deal evenly' % num_players)
        hands = []
        for i, card in enumerate(self.cards):
            hand_len = i % num_players
            if len(hands) < hand_len + 1:
                hands.append([])
            hands[hand_len].append(card)
        self.hands = hands


def main():
    print "Let's play war!"

    # Create the deck
    deck = Deck(RANK_DICT, SUIT_LIST)
    deck.populate()
    deck.shuffle()
    deck.deal(2)

    # Play the game!
    num_turns = 0
    num_war = 0
    while(all([len(hand) for hand in deck.hands])):
        winner, pile, n_war = compare_cards(*deck.hands)
        num_war += n_war

        # Shuffle pile and add to winners hand
        random.shuffle(pile)
        deck.hands[winner].extend(pile)

        if PRINT_HANDS:
            print [len(hand) for hand in deck.hands]
        num_turns += 1
        if num_turns > MAX_TURNS:
            print "Too many turns, quitting game"
            break

    # Print winner and stats
    for i, hand in enumerate(deck.hands):
        if len(hand):
            winner = i + 1
    print "Congrats player %d" % winner
    print "Number of turns: %d" % num_turns
    print "Number of wars: %d" % num_war


def compare_cards(deck1, deck2):
    """
    Compare cards from the deck

    This method takes two decks and returns a winner
    along with a pile of cards the winner gets to keep
    """
    c1 = deck1.pop(0)
    c2 = deck2.pop(0)

    winner = 0
    pile = [c1, c2]
    num_war = 0
    if c1 > c2:
        winner = 0
    if c1 < c2:
        winner = 1
    if c1 == c2:
        print '\tWar!'
        num_war += 1
        for deck in (deck1, deck2):
            pile.extend(war_pile(deck))

        # Always compare at least the last exposed card
        if len(deck1) == 0:
            deck1.append(c1)
        if len(deck2) == 0:
            deck2.append(c2)

        # Compare the cards and see who wins
        winner, pile2, n_war = compare_cards(deck1, deck2)
        pile.extend(pile2)
        num_war += n_war

    # Return the winner and the pile
    return winner, list(set(pile)), num_war


def war_pile(deck):
    """ Get up to 3 cards from the deck """
    pile = []
    for i in range(0, 3):
        if len(deck) > 1:
            pile.append(deck.pop(0))
    return pile


if __name__ == "__main__":
    main()
