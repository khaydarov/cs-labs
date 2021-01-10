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

(define (prime-sum? pair)
	(prime? (+ (car pair) (cadr pair))))

(define (make-pair-sum pair)
	(list (car pair) (cadr pair) (+ (car pair) (cadr pair))))

(define (unique-pairs n)
	(flatMap
		(lambda (i)
			(map (lambda (p) (cons i j)) (enum-interval 1 (- i 1))))
		(enum-interval 1 n)))

(define (prime-sum-pairs n)
	(map make-pair-sum
		(filter prime-sum? (unique-pairs n))))
