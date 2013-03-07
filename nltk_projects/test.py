#!/usr/local/bin/python

import cPickle
import nltk
import os


class Ingredient(object):
    name = None
    amount = None
    unit = None
    grams = []


def main():
    filename = os.path.abspath(os.path.join(os.getcwd(), 't1.pkl'))
    ingredients = [
            "Golden raisins (1/3 cup)",
            "Green onions (1 cup)",
            "Lemon, zest (1 teaspoon)",
            "Pine nuts (1/2 cup)",
            "Balsamic vinegar, white (2 tablespoons)",
            "Olive oil, extra virgin (2 tablespoons)",
            "Pepper, freshly ground (1/8 teaspoon)",
            "Salt (3/4 teaspoon)",
            "Carolina gold long-grain rice (1 1/2 cups)",
            "Lemon juice, fresh (1/4 cup)",
            "Orange juice, fresh (2 tablespoons)",
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
            (r'.*', 'NN')  # Noun tagger
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

    # Train bigram tagger based on tagged sentences
    t2 = nltk.BigramTagger(dir_tagged_sents, backoff=t1)

    # Make some useful sets of tagged words
    unigram_directions = [unigram for sent in dir_tagged_sents for unigram in sent]
    bigram_directions = [bigram for sent in dir_tagged_sents for bigram in nltk.bigrams(sent)]

    # Work through the ingredients and compare them to directions
    saved_ingredients = []
    extra_ingredients = []
    for ing in ingredients:
        found = False
        text = nltk.word_tokenize(ing.lower().replace(",", ""))
        text = [word for word in text if text not in nltk.corpus.stopwords.words('english')]
        tagged = t2.tag(text)
        bigrams = nltk.bigrams(tagged)

        for bi in bigrams:
            if bi in bigram_directions:
                found = True
                saved_ingredients.append(bi)
            else:
                for tag in bi:
                    exempt = [t for b in saved_ingredients for t in b]
                    if tag in unigram_directions and 'NN' in tag[1] and tag not in exempt:
                        found = True
                        saved_ingredients.append([tag])
        if not found:
            extra_ingredients.append(tagged)

    # Work through the extra ingredients for things not in directions
    for ing in extra_ingredients:
        new_list = []
        found = False
        for tag in ing:
            if 'NN' in tag[1] or 'JJ' in tag[1]:
                found = True
                new_list.append(tag)
            elif found:
                break
        saved_ingredients.append(new_list)

    # Print out the final ingredients
    for ing in saved_ingredients:
        sentence = []
        found = False
        for b in ing:
            if 'NN' in b[1] or ('NN' not in b[1] and not found):
                sentence.append(b[0])
            if 'NN' in b[1]:
                found = True

        print ' '.join(sentence)


if __name__ == "__main__":
    main()
