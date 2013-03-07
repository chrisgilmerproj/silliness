#!/usr/local/bin/python

import cPickle
import nltk
import os

UNITS = [
    "teaspoon", "teaspoons", "t", "tsp",
    "tablespoon", "tablespoons", "T", "tbl", "tbs", "tbsp",
    "ounce", "oz",
    "gill", "gills",
    "cup", "cups", "c",
    "pint", "pints", "p", "pt",
    "quart", "quarts", "q", "qt",
    "gallon", "gallons", "g", "gal",
]

class Ingredient(object):

    def __init__(self, text, ordinal=None, tagger=None):
        self.text = text.lower()
        self.tokens = nltk.word_tokenize(self.text.lower())
        self.unigrams = tagger.tag(self.tokens)
        self.bigrams = nltk.bigrams(self.unigrams)
        self.tagged = []

        self.ordinal = ordinal
        self.amount = None
        self.unit = None
        self.name = None
        self.modifier = None

    def __repr__(self):
        label = "{0:2}. {1} {2} {3}".format(self.ordinal, self.amount, self.unit, self.name)
        if self.modifier:
            label += ", {0}".format(self.modifier)
        return label


def main():
    filename = os.path.abspath(os.path.join(os.getcwd(), 't1.pkl'))
    ingredients = [
            "Pine nuts (1/2 cup)",
            "Carolina gold long-grain rice (1 1/2 cups)",
            "Lemon, zest (1 teaspoon)",
            "Lemon juice, fresh (1/4 cup)",
            "Orange juice, fresh (2 tablespoons)",
            "Balsamic vinegar, white (2 tablespoons)",
            "Salt (3/4 teaspoon)",
            "Pepper, freshly ground (1/8 teaspoon)",
            "Olive oil, extra virgin (2 tablespoons)",
            "Green onions (1 cup)",
            "Golden raisins (1/3 cup)",
            ]
    ingredients = [
            "1/2 cup pine nuts",
            "1 1/2 cups uncooked Carolina Gold long-grain rice",
            "1 teaspoon lemon zest",
            "1/4 cup fresh lemon juice (about 2 lemons)",
            "2 tablespoons fresh orange juice",
            "2 tablespoons white balsamic vinegar",
            "3/4 teaspoon salt",
            "1/8 teaspoon freshly ground pepper",
            "2 tablespoons extra virgin olive oil",
            "1 cup sliced green onions",
            "1/3 cup golden raisins",
            ]
    directions = """Cook pine nuts in a skillet over medium-low heat, stirring often, 5 minutes or until toasted and fragrant.
    Cook rice according to package directions. 
    (You should have about 6 cups.)
    Whisk together lemon zest and next 5 ingredients. 
    Add olive oil in a slow, steady stream, whisking until blended and smooth.
    Combine hot cooked rice, green onions, raisins, and pine nuts. 
    Pour dressing over rice mixture, and stir until blended. 
    Let stand 30 minutes before serving, or cover, chill, and serve the next day."""

    # Save everything the first time and load it later
    if os.path.isfile(filename):
        print 'Loading saved tagger from {0}'.format(filename)
        with open(filename, 'rb') as f:
            t1 = cPickle.load(f)
    else:
        patterns = [
            (r'^\d\/\d$', 'CD'),  # Fractions
            (r'.*\d.*', 'CD'),    # More numbers
            (r'.*ed$', 'VBD'),    # simple past
            (r'.*', 'NN'),        # Noun tagger
        ]
        train_sents = nltk.corpus.brown.tagged_sents()
        t0 = nltk.RegexpTagger(patterns)
        t1 = nltk.UnigramTagger(train_sents, backoff=t0)
        with open(filename, 'wb') as f:
            cPickle.dump(t1, f, -1)

    # Tag the directions using the unigram tagger
    dir_tagged_sents = []
    for sentence in nltk.tokenize.sent_tokenize(directions):
        text = nltk.word_tokenize(sentence.rstrip('.'))
        tagged = t1.tag(text)
        dir_tagged_sents.append(tagged)

    # Train bigram tagger based on tagged direction sentences
    t2 = nltk.BigramTagger(dir_tagged_sents, backoff=t1)

    # Make some useful sets of tagged words
    unigram_directions = [unigram for sent in dir_tagged_sents
                                  for unigram in sent]
    bigram_directions = [bigram for sent in dir_tagged_sents
                                for bigram in nltk.bigrams(sent)]

    # Work through the ingredients and compare them to directions
    ingredient_list = []
    for num, sentence in enumerate(ingredients, start=1):
        ingredient = Ingredient(sentence, ordinal=num, tagger=t2)
        ingredient_list.append(ingredient)

    # First pass looks at the directions and tries to match bigrams and unigrams
    exempt = []  # Once items identified in directions they cannot be identified again
    for ingredient in ingredient_list:
        # Find bigrams in directions first
        for bigram in ingredient.bigrams:
            if bigram in bigram_directions and \
                    bigram not in exempt:
                ingredient.tagged = list(bigram)
                exempt.extend(list(bigram))

        # If bigrams not found then find unigrams in directions minus exempt
        if not ingredient.tagged:
            for unigram in ingredient.unigrams:
                if unigram in unigram_directions and \
                        'NN' in unigram[1] and \
                        unigram not in exempt and \
                        unigram[0] not in UNITS:
                    exempt.append(unigram)
                    ingredient.tagged.append(unigram)

    # Work through the ingredients that were not tagged
    for ingredient in filter(lambda x: not x.tagged, ingredient_list):
        last = ''
        found = False
        for tag in ingredient.unigrams:
            if 'NN' in tag[1] and last != 'RB' and tag[0] not in UNITS:
                found = True
                ingredient.tagged.append(tag)
            elif found:
                break
            last = tag[1]

    # Set the name for each ingredient
    for ingredient in ingredient_list:

        # Put together the ingredient name
        name = []
        found = False
        for bigram in ingredient.tagged:
            if 'NN' in bigram[1]:
                name.append(bigram[0])
                found = True
            if 'JJ' in bigram[1] and not found:
                name.append(bigram[0])
        ingredient.name = ' '.join(name)

        # Put together the amount and unit
        amount = []
        last = ''
        for unigram in ingredient.unigrams:
            if 'CD' in unigram[1]:
                amount.append(unigram[0])
            if 'NN' in unigram[1] and last == 'CD':
                ingredient.unit = unigram[0]
                break
            last = unigram[1]
        ingredient.amount = ' '.join(amount)

        modifier = []
        for unigram in ingredient.unigrams:
            if unigram[0] not in name + amount + [ingredient.unit]:
                modifier.append(unigram[0])
        ingredient.modifier = " ".join(modifier)

    # Print out everything
    for ingredient in ingredient_list:
        print "{0:75} = {1}".format(ingredient, ingredient.text)


if __name__ == "__main__":
    main()
