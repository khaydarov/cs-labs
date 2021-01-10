(define nil '())

(define (enum-interval low high)
	(if (> low high)
		nil
		(cons low (enum-interval (+ low 1) high))))

(define (accumulate op initial sequence)
	(if (null? sequence)
		initial
		(op (car sequence)
			(accumulate op initial (cdr sequence)))))

(define (flatMap proc seq)
	(accumulate append nil (map proc seq)))

(define (triple-sum triple)
	(+ (car triple) (cadr triple) (cadr (cdr triple))))

(define (unique-triples n)
	(flatMap
		(lambda (i)
			(flatMap
				(lambda (j)
					(map (lambda (k) (list i j k)) (enum-interval 1 (- j 1))))
			(enum-interval 1 (- i 1))))
		(enum-interval 1 n)))

(define (make-triple-sum n s)
	(filter (lambda (triple)
				(equal? (triple-sum triple) s))
			(unique-triples n)))
