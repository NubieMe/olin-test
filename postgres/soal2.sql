select
    u.name, 
    o.amount, 
    o.created_at 
from 
    first.public.users as u
join 
    second.public.orders as o on u.id = o.user_id;