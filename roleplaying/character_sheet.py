#! /usr/local/bin/python

import argparse
import json
from pprint import pprint
import string
import sys


def abilities(character):
    print '\nABILITIES AND SKILLS'
    for ability in character['abilities and skills']:
        name = ability['name']
        value = ability['value']
        modifier = ability['modifier']
        abilities = ability['abilities']
        print "(%s) %s +%d" % (value, name, modifier)
        for ab, desc in abilities.iteritems():
            print '\t%s%s%s' % (string.ljust(ab, 15), string.ljust(str(desc['trained']), 6), desc['value'])


def combat(character):
    print '\nCOMBAT STATISTICS'
    data = character['combat statistics']
    print '(%s) Initiative\t(%s) Speed' % (data['initiative'], data['speed'])
    print '\nDEFENSES'
    for item in data['defenses']:
        print '\t(%s) %s' % (item['value'], item['name'])
    print '\nWEAPONS'
    for item in data['weapons']:
        print '\t(%s) %s %s' % (item['value'], item['name'], item['damage'])
    print '\nHEALTH'
    hit_points = data['hit points']
    print '\t(%s) Hit Points' % (hit_points['value'])
    print '\t(%s) Bloodied' % (hit_points['bloodied'])
    print '\tHealing Surges: %s' % (hit_points['healing surge value'])
    print '\tSurges per Day: %s' % (hit_points['surges per day'])
    print '\nCURRENT HIT POINTS'
    print '\tTemporary Hit Pints: %s' % (data['current hit points']['temporary hit points'])
    print '\tSurges Used: %s' % (data['current hit points']['surges used'])


def equipment(character):
    print "\nEQUIPMENT AND MAGIC ITEMS"
    for item in character['equipment and magic items']:
        print '\t%s' % item


def notes(character):
    print "\nCHARACTER NOTES"
    for item in character['character notes']:
        print '\t%s' % item


def powers(character):
    print "\nPOWERS AND FEATS"
    for item in character['powers and feats']:
        print '\t%s' % item


def wealth(character):
    print "\nWEALTH"
    data = character['wealth']
    print "\t%s gold" % (string.rjust(str(data['gold']), 5))
    print "\t%s silver" % (string.rjust(str(data['silver']), 5))
    print "\t%s copper" % (string.rjust(str(data['copper']), 5))
    print "\t%s" % (data['notes'])


def experience(character):
    print "\nEXPERIENCE POINTS (XP)"
    data = character['experience points (xp)']
    value = float(data['value'])
    level_up = float(data['level up'])
    next = level_up - value
    print "\t%s current points" % (value)
    print "\t%s next level up" % (level_up)
    print "\t%s remaining" % (next)


def main(args):

    try:
        json_data = open(args.filename, 'r').read()
    except (IOError, TypeError):
        print "Cannot load the file '%s'" % (args.filename)
        sys.exit()

    try:
        character = json.loads(json_data)
    except ValueError:
        print "Your file '%s' is improperly formatted or missing" % (args.filename)
        sys.exit()

    person = character['character']
    print '\nCHARACTER'
    print '\t%s: %s' % (string.ljust('Name', 10), person['name'])
    print '\t%s: %s' % (string.ljust('Class', 10), person['class'])
    print '\t%s: %s' % (string.ljust('Level', 10), person['level'])
    print '\t%s: %s' % (string.ljust('Race', 10), person['race'])
    print '\t%s: %s' % (string.ljust('Gender', 10), person['gender'])
    print '\t%s: %s' % (string.ljust('Alignment', 10), person['alignment'])
    print '\t%s: %s' % (string.ljust('Languages', 10), person['languages'])

    if args.abilities:
        abilities(character)
    if args.combat:
        combat(character)
    if args.equipment:
        equipment(character)
    if args.notes:
        notes(character)
    if args.powers:
        powers(character)
    if args.wealth:
        wealth(character)
    if args.experience:
        experience(character)


if __name__ == "__main__":

    parser = argparse.ArgumentParser(description='Character Sheet Program')
    parser.add_argument('-f', '--file', dest='filename', action='store',
                   help='Load a character sheet')
    parser.add_argument('-a', '--abilities', dest='abilities', action='store_true',
                   default=False, help='Show Abilities and Skills')
    parser.add_argument('-c', '--combat', dest='combat', action='store_true',
                   default=False, help='Show Combat Statistics')
    parser.add_argument('-e', '--equipment', dest='equipment', action='store_true',
                   default=False, help='Show Equipment and Magic Items')
    parser.add_argument('-n', '--notes', dest='notes', action='store_true',
                   default=False, help='Show Character Notes')
    parser.add_argument('-p', '--powers', dest='powers', action='store_true',
                   default=False, help='Show Powers and Feats')
    parser.add_argument('-w', '--wealth', dest='wealth', action='store_true',
                   default=False, help='Show Wealth')
    parser.add_argument('-x', '--experience', dest='experience', action='store_true',
                   default=False, help='Show Experience Points (XP)')

    args = parser.parse_args()
    main(args)
