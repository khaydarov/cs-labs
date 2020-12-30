(define (make-interval a b)
	(if (< a b)
		(cons a b)
		(cons b a)))

(define (lower-bound i)
	(car i))

(define (upper-bound i)
	(cdr i))

(define (mcp center perc)
	(make-interval (- center (/ (* perc center) 100))
					    	 (+ center (/ (* perc center) 100))))

 (define (perc-selector i)
	(define (center i)
		(/ (+ (lower-bound i) (upper-bound i)) 2))
	(let ((c (center i)))
		(/ (* (- c (lower-bound i)) 100) c)))
