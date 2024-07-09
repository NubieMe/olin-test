select
    u.name, sum(o.price * o.quantity) as total
from
    users as u
join
    orders as o on u.id = o.user_id
join
    order_items as i on i.order_id = o.id
where
    o.created_at >= '2020-01-01'
group by
    u.name
having
    total >= 100.00;