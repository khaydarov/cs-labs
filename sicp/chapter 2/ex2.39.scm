(define nil '())

;Reverse in terms of fold-right
(define (rev-fold-right seq)
	(fold-right (lambda (x y)
					(if (null? y)
							(list x)
							(append y (list x))))
				nil
				seq))

;Reverse in terms of fold-left
(define (rev-fold-left seq)
	(fold-left (lambda (x y) (cons y x)) nil seq))
