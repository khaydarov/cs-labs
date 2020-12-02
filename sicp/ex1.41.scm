(define (double f)
	(lambda (x) (f (f x))))

;Returns 3
((double inc) 1)

;Returns 21
(((double (double double)) inc) 5)
