#! /usr/local/bin/python
import sys
import time

from googlevoice import Voice
from googlevoice.util import input

TIMEOUT = 60

if __name__ == "__main__":

    print '\nInitiating Google Voice Connection'

    voice = Voice()
    voice.login()
    
    #TODO: Use regular expression to get number
    phone_number = input('Enter a phone number for the session (XXXXXXXXXX): ')

    open_connection = time.time()
    
    text = None
    if len(sys.argv) == 2:
        text = str(sys.argv.pop())
    
    try:
        while 1:
            while not text:
                text = input('\nMessage text: ')
            
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
    except KeyboardInterrupt, e:
        print '\n\nThank you for using Google Voice'
