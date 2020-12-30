(define zero (lambda (f) (lambda (x) x)))

(define (add-1 n)
	(lambda (f)
		(lambda (x) (f ((n f) x)))))

(define one (lambda (f) (lambda (f x))))
(define two (lambda (f) (lambda (f (f x)))))

(define (add n m)
	(lambda (f)
		(lambda (x) ((n f) ((m f) x)))))
