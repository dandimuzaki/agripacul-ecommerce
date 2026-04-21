'use client';

import { Button } from '@/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Spinner } from '@/components/ui/spinner';
import useLogin from '@/hooks/auth/useLogin';
import { LoginFormValues, loginSchema } from '@/schemas/auth.schema';
import { zodResolver } from '@hookform/resolvers/zod';
import { Eye, EyeOff } from 'lucide-react';
import Image from 'next/image';
import { useState } from 'react'
import { useForm } from 'react-hook-form';

const LoginCustomer = () => {
  const [isSeenPassword, setIsSeenPassword] = useState(false)

  const values: LoginFormValues = {
    email: "",
    password: ""
  }

  const form = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: values,
  })

  const { mutate: login, isPending, error, isError } = useLogin()
  
  const onSubmit = (data: LoginFormValues) => {
    login(data)
  }

  return (
    <div className='h-screen grid lg:grid-cols-2'>
      <div className='bg-primary'>
      <div className='relative hidden lg:block lg:h-full overflow-hidden'>
        <img 
          src="/vegetable-garden.jpg" 
          alt="Salad Shake by Agripacul" 
          width={100} height={100} 
          className="h-full w-full object-cover"
          loading="eager"
           />
          <div className="p-12 flex flex-col gap-2 justify-end absolute z-2 bottom-0 h-full w-full bg-[linear-gradient(rgba(0,0,0,0),rgba(0,0,0,0.8))]">
            <p className="text-3xl text-white font-bold">Welcome back to Agripacul</p>
            <p className="text-lg/6 text-white">We connect you with fresh produce, homemade foods, and gardening essentials, grown responsibly and delivered with care for everyday cooking and growing</p>
          </div>
      </div>
      </div>
    <div className='px-8 md:px-16 bg-primary flex flex-col gap-6 items-center justify-center'>
      <div className='h-10'>
        <Image src="/logo.png" alt="Logo Agripacul" height={100} width={100} className='w-full h-full object-cover' />
      </div>
      <Card className="w-full max-w-sm space-y-4">
        <CardHeader>
          <CardTitle>Login to your account</CardTitle>
        </CardHeader>
        <CardContent>
          <form 
            id='login-customer'
            onSubmit={form.handleSubmit(onSubmit)}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="fulan@example.com"
                  required
                  {...form.register("email")}
                />
              </div>
              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                  <a
                    href="#"
                    className="ml-auto inline-block text-sm underline-offset-4 hover:underline"
                  >
                    Forgot your password?
                  </a>
                </div>
                <div className='relative'>
                  <Input 
                    id="password" 
                    type={`${isSeenPassword ? "text" : "password" }`}
                    required
                    {...form.register("password")}
                  />
                  <div 
                    onClick={() => setIsSeenPassword(prev => !prev)}
                    className='absolute bottom-2 right-2 text-lg'
                  >
                    {isSeenPassword ? <EyeOff fontSize="inhert"/> : <Eye fontSize="inhert"/>}
                  </div>
                 </div>
              </div>
            </div>
          </form>
        </CardContent>
        <CardFooter className="flex-col gap-2">
          <Button form='login-customer' type="submit" className="w-full text-base mt-2">
            {isPending ? <><Spinner/>Logging in...</> : <>Login</>}
          </Button>
          {isError && <p className='text-red-500 text-sm'>Something went wrong: {error.message}</p>}
        </CardFooter>
      </Card>
    </div>
    </div>
  )
}

export default LoginCustomer
