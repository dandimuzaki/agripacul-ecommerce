'use client';

import { Button } from '@/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Spinner } from '@/components/ui/spinner';
import { LoginFormValues, loginSchema } from '@/schemas/auth.schema';
import { authService } from '@/services/auth.service';
import { zodResolver } from '@hookform/resolvers/zod';
import { Eye, EyeOff } from 'lucide-react';
import Image from 'next/image';
import { useRouter } from 'next/navigation';
import { useState } from 'react'
import { useForm } from 'react-hook-form';

const LoginAdmin = () => {
  const router = useRouter()
  const [isSeenPassword, setIsSeenPassword] = useState(false)

  const values: LoginFormValues = {
    email: "",
    password: ""
  }

  const form = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: values,
  })
  
  const [isLoading, setIsLoading] = useState(false)

  const onSubmit = async (data: LoginFormValues) => {
    try {
      setIsLoading(true)
      await authService.loginAdmin(data)
      router.push("/admin/products")
    } catch (error) {
      console.error(error)
    } finally {
      setIsLoading(false)
    }
  }  

  return (
    <div className='h-screen bg-primary flex flex-col gap-6 items-center justify-center'>
      <div className='h-10'>
        <Image src="/logo.png" alt="Logo Agripacul" height={100} width={100} className='w-full h-full object-cover' />
      </div>
      <Card className="w-full max-w-sm space-y-4">
        <CardHeader>
          <CardTitle>Login to your account</CardTitle>
        </CardHeader>
        <CardContent>
          <form id="login-admin" onSubmit={form.handleSubmit(onSubmit)}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="m@example.com"
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
                    className='absolute bottom-2 right-2'
                  >
                    {isSeenPassword ? <EyeOff/> : <Eye/>}
                  </div>
                 </div>
              </div>
            </div>
          </form>
        </CardContent>
        <CardFooter className="flex-col gap-2">
          <Button form="login-admin" type="submit" className="w-full">
            {isLoading ? <><Spinner/>Logging in...</> : <>Login</>}
          </Button>
        </CardFooter>
      </Card>
    </div>
  )
}

export default LoginAdmin
