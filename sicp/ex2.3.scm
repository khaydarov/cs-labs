(define (make-point a b)
	(cons a b))

(define (x-point p)
	(car p))

(define (y-point p)
	(cdr p))

(define (make-rect bottom-left top-right)
	(cons bottom-left top-right))

(define (bottom-right rect)
	(make-point (x-point (cdr rect))
							(y-point (car rect))))

(define (top-left rect)
	(make-point (x-point (car rect))
						  (y-point (cdr rect))))

(define (width rect)
	(let ((x2 (x-point (bottom-right rect)))
				(x1 (x-point (car rect))))
				(- x2 x1)))

(define (height rect)
	(let ((y2 (y-point (top-left rect)))
				(y1 (y-point (car rect))))
				(- y2 y1)))

(define (perimeter rect)
	(let ((w (width rect))
			  (h (height rect)))
			(* (+ w h) 2)))

(define (area rect)
	(let ((w (width rect))
			  (h (height rect)))
			(* w h)))
