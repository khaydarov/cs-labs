;Iterative implementation
(define (repeated f n)
	(define (iter n result)
		(if (= n 0)
				result
				(iter (- n 1) (f result))))

	(lambda (x) (iter n x)))

;Recursive implementation
(define (repeated f n)
	(define (compose f g)
		(lambda (x) (f (g x))))

	(if (> n 1)
			(lambda (x) ((compose (repeated f (- n 1)) f) x))
			f))
