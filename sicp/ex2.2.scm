(define (make-point a b)
	(cons a b))

(define (x-point p)
	(car p))

(define (y-point p)
	(cdr p))

(define (make-segment a b)
	(cons a b))

(define (start-segment s)
	(car s))

(define (end-segment s)
	(cdr s))

(define (print-point p)
		(newline)
		(display "(")
		(display (x-point p))
		(display ",")
		(display (y-point p))
		(display ")"))

(define (average a b)
	(/ (+ a b) 2))

(define (midpoint-segment s)
		(let ((start (start-segment s))
			    (end (end-segment s)))

		(make-point (average (x-point start)
												 (x-point end))
								(average (y-point start)
												 (y-point end)))))
