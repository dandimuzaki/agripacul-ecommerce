'use client'

import Link from 'next/link'
import LinkSidebar from './link-sidebar'
import Image from 'next/image'
import useLogout from '@/hooks/auth/useLogout'

const Sidebar = () => {
  const paths = [
    {
      path: "/admin/category",
      name: "Category",
      icon: "lucide:chart-pie"
    },
    {
      path: "/admin/product",
      name: "Product",
      icon: "lucide:salad"
    },
    {
      path: "/admin/order",
      name: "Order",
      icon: "lucide:shopping-bag"
    },
    {
      path: "/admin/inventory",
      name: "Inventory",
      icon: "lucide:shelving-unit"
    },
    {
      path: "/admin/review",
      name: "Review",
      icon: "lucide:message-square-text"
    },
  ]

  const {mutate: onLogout} = useLogout()
  
  return (
    <div className='fixed flex flex-col top-0 w-64 h-screen bg-primary text-white'>
      <div className='h-10 w-fit flex justify-center px-6 my-3'>
        <Link href="/admin" className='w-fit'>
          <Image src="/logo.png" height={100} width={100} alt='logo Agripacul' className='object-cover h-8 w-full' />
        </Link>
      </div>
      {paths.map((p, i) => (
        <LinkSidebar
          key={i}
          path={p.path}
          name={p.name}
          icon={p.icon}
        />
      ))}
      <div className='cursor-pointer flex-1 flex items-end justify-center'>
      <p onClick={() => onLogout()} className='px-6 py-3 text-lg hover:bg-primary-foreground w-full'>Logout</p>
      </div>
    </div>
  )
}

export default Sidebar
