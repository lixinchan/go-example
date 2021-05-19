# -*- coding: utf-8 -*-

from french_deck import FrenchDeck
from french_deck import Card
from random import choice

if __name__ == '__main__':
    deck = FrenchDeck()
    print(deck.ranks)
    print(deck.suits)
    print(deck[0])
    print(deck[-1])
    print(choice(deck))
    print(len(deck))
    # slice
    print(deck[:3])
    print(deck[12::13])
    # loop
    for card in deck:
        print(card)
    for card in reversed(deck):
        print(card)
    # contains
    print(Card('10', 'spades') in deck)
    # sort
    for card in sorted(deck, key=deck.spades_high):
        print(card)
