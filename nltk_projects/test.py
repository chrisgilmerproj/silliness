#!/usr/local/bin/python

import nltk


def main():
    sentence_list = [
        "And now for something completely different",
        "They refuse to permit us to obtain the refuse permit",
        ]
    for sentence in sentence_list:
        text = nltk.word_tokenize(sentence)
        tagged = nltk.pos_tag(text)
        print sentence
        print text
        print tagged


if __name__ == "__main__":
    main()
