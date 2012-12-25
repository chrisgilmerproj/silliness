#! /usr/local/bin/python

import os
import random
import string
import sys


class CardsAgainstHumanity():
    """
    http://www.cardsagainsthumanity.com/

    Please read LICENSE.  Current license can be found at the website above.
    """

    def __init__(self):
        self.white_card_names = ["wcards.txt"]
        self.white_cards = []

        self.blank = "__________"
        self.black_card_names = ["bcards.txt", "bcards1.txt", "bcards2.txt"]
        self.black_cards = []

        self.hand = []
        self.hand_size = 7

        self.setup()

    def setup(self):
        # Get the cards
        self.white_cards = self.get_cards(self.white_card_names)
        self.black_cards = self.get_cards(self.black_card_names)

        # Shuffle the cards
        random.shuffle(self.white_cards)
        random.shuffle(self.black_cards)

        self.deal_hand()

    def get_cards(self, card_names):
        card_list = []
        for name in card_names:
            with open(os.path.join("data", name)) as f:
                for line in f:
                    cards = line.strip().split("<>")
                    for card in cards:
                        card_list.append(card.strip('.'))
        card_list = list(set(card_list))
        return card_list

    def draw_card(self):
        self.hand.append(self.white_cards.pop())

    def deal_hand(self):
        while len(self.hand) != self.hand_size:
            self.draw_card()

    def get_required_cards(self, question):
        req = string.count(question, self.blank)
        if req == 0:
            req = 1
        return req

    def get_question_card(self):
        return self.black_cards.pop()

    def get_answer_card(self, number):
        return self.hand.pop(number - 1)

    def format_question_and_answer(self, question, answers):
        if len(answers) == 1 and self.blank not in question:
            return "  {0} - '{1}'".format(question, " ".join(answers))
        else:
            count = 1
            for a in answers:
                question = question.replace(self.blank,
                                            "'{0}'".format(a), count)
            return "  {0}".format(question)

    def run(self):

        while len(self.black_cards):

            question = self.get_question_card()
            num_required = self.get_required_cards(question)

            # Deal extra cards to player
            if num_required > 1:
                for x in xrange(num_required - 1):
                    self.draw_card()

            print '=' * 80
            print
            print question
            print

            number = 0
            answers = []
            while True:
                for num, card in enumerate(self.hand, start=1):
                    print num, card

                try:
                    print
                    message = "Choose a card to play: "
                    number = int(raw_input(message))
                except ValueError:
                    print "Invalid choice, choose again ..."

                if 0 < number < len(self.hand) + 1:
                    answers.append(self.get_answer_card(number))
                    if len(answers) == num_required:
                        self.draw_card()
                        print
                        print "=" * 80
                        print self.format_question_and_answer(question, answers)
                        print "=" * 80
                        print
                        val = raw_input("Press enter to continue game ... ")
                        if val in ['q', 'Q']:
                            sys.exit()
                        print
                        break


if __name__ == "__main__":
    cah = CardsAgainstHumanity()
    try:
        cah.run()
    except:
        sys.exit()
