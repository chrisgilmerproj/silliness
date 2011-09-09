#! /usr/local/bin/python
import re
import sys
import time

from googlevoice import Voice
from googlevoice.util import input

TIMEOUT = 60

REXP_PHONE = re.compile('''(\d{3})\D*(\d{3})\D*(\d{4})''')

if __name__ == "__main__":

    print '\nInitiating Google Voice Connection'

    voice = Voice()
    voice.login()
    
    phone_number = ''
    while phone_number == '':
        num = input('Enter a phone number for the session (XXX-XXX-XXXX): ')
        phone_rexp = REXP_PHONE.search(num)
        if phone_rexp:
            phone_number = ''.join(phone_rexp.groups())
    print '\nUsing: %s' % phone_number

    open_connection = time.time()
    
    text = None
    if len(sys.argv) == 2:
        text = str(sys.argv.pop())
    
    while 1:
        # It's common to exit with ^C, so catch it
        try:
            while not text:
                text = input('\nMessage text: ')
        except KeyboardInterrupt, e:
            print '\n\nThank you for using Google Voice'
            sys.exit(1)
        
        if time.time() - open_connection > TIMEOUT:
            print '\nConnection timed out after %s seconds' % TIMEOUT
            break

        print '%s: %s' % (phone_number, text)

        send = raw_input('Would you like to send? y/N ')
        if send in ['y','Y','yes','Yes','YES']:
            print 'Sending message'
            voice.send_sms(phone_number, text)
        else:
            print 'Aborting message'

        text = None
