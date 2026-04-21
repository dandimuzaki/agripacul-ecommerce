import Link from 'next/link'
import { usePathname } from 'next/navigation';
import { Icon } from '../ui/icon';

export default function LinkSidebar({
  path, 
  name, 
  icon}: 
  ({
    path: string, 
    name: string, 
    icon: string
  })) {
  
  const pathname = usePathname()
  const isActive = (path: string) => pathname === path
  const baseClass =
    'px-6 py-3 flex gap-2 items-center transition-colors text-lg'

  const activeClass = 'bg-primary-foreground font-bold'
  const inactiveClass = 'bg-primary'

  return (
    <Link
        href={path}
        className={`${baseClass} ${
          isActive(path) ? activeClass : inactiveClass
        }`}
      >
        <Icon icon={icon} /> {name}
      </Link>
  );
};