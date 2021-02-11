(define (adjoin-set x set)
	(if (> x (car set))
		(cons (car set) (adjoin-set x (cdr set)))
		(cons x set)))
