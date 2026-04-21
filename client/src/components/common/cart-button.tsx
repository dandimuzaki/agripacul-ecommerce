import { useCart } from '@/hooks/cart/useCart';
import { ShoppingCartOutlined } from '@mui/icons-material';
import Link from 'next/link';
import { Button } from '../ui/button';

const CartButton = ({isHome, isOpen, scrolled}: {isHome: boolean, isOpen: boolean, scrolled: boolean}) => {
  const { data: cart } = useCart();
  return (
    <Link href='/cart'>
      <Button className={`relative h-full ${isHome && !isOpen && !scrolled && "bg-transparent text-primary hover:text-white"}`}>
        <ShoppingCartOutlined />
        <span className='counter absolute top-0 right-0 text-white flex bg-red-500 text-xs h-[18px] items-center justify-center rounded-full aspect-square'>{cart?.summary.total_items || 0}</span>
      </Button>
    </Link>
  );
};

export default CartButton;