(define nil '())
(define (accumulate op initial sequence)
	(if (null? sequence)
		initial
		(op (car sequence)
			(accumulate op initial (cdr sequence)))))

(define (accumulate-n op init seqs)
	(if (null? (car seqs))
		nil
		(cons (accumulate op init (map car seqs))
			  (accumulate-n op init (map cdr seqs)))))

(define m (list (list 1 2 3 4) (list 4 5 6 6) (list 6 7 8 9)))
(define v (list 1 2 3 4))

(define (dot-product v w)
	(map * v w)

; matrix-*-vector implementation
(define (matrix-*-vector m v)
	(if (null? m)
		nil
		(cons (accumulate + 0 (map * (car m) v))
				  (matrix-*-vector (cdr m) v))))
; or
(define (matrix-*-vector m v)
	(map
		(lambda (row)
			(accumulate + 0 (map * row v)))
		m))

; Transpose implementation
(define (transpose mat)
	(accumulate-n cons nil mat))

; Matrix to martix multiplication
(define (matrix-*-matrix m n)
	(let ((cols (transpose n)))
		(map (lambda (row)
			(map (lambda (col)
				(accumulate + 0 (map * row col)))
				cols))
			m)))

; or
(define (matrix-*-matrix m n)
	(let ((cols (transpose n)))
		(map (lambda (row) (matrix-*-vector cols row)) m)))
