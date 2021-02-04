(define (equal? x y)
	(cond  ((and (null? x) (null? y)) #t)
		   ((or  (and (null? x) (not (null? y)))
			     (and (not (null? x)) (null? y))) #f)
		   ((not (eq? (car x) (car y))) #f)
		   (else (equal? (cdr x) (cdr y)))))
