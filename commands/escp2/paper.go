package escp2

// SetDefaultUnit sets measuring unit to x / 3600 inch, however different settings may alter the currently used
// unit in the printer.
// For example:
//
//			using the move to horizontally / vertically sets this value to 10
//	     	using letter quality (LQ) : 20 / 3600" (1/180")
//
// When passing improper value this command resets to standard 1/360"
