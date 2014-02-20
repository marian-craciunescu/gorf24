/*  Copyright 2013, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gorf24>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package xaddr

/* used for Rx and Tx addresses; hence 'x' address
   max 5 bytes
   using 64bits for simple storage, but will truncate
   first 3 bytes to limit to 5 bytes

   64bits
   <-- ignored -----------><-- used ------------------------------>
   xxxxxxxxxxxxxxxxxxxxxxxx1111111111111111111111111111111111111111 */
type A uint64

const RES_MASK uint64 = 0x000000FFFFFFFFFF

func New(flags uint64) A {
	// mask out first 3 bytes
	return A(flags & RES_MASK)
}

func (a A) Byte() uint64 {
	return uint64(a)
}
