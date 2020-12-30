;Recursive implementation
(define (repeated f n)
	(define (compose f g)
		(lambda (x) (f (g x))))

	(if (> n 1)
			(lambda (x) ((compose (repeated f (- n 1)) f) x))
			f))

;Smooth implementation
(define (smooth f)
	(define dx 0.0001)
		(lambda (x)
			(/ (+ (f (- x dx)) (f x) (f (+ x dx))) 3)))

;n-smooth implementation
(define (n-smooth f n)
	(lambda (x) ((repeated (smooth f) n) x)))
