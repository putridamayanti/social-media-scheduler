'use client'

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import {useRouter} from "next/navigation";
import {useFormik} from "formik";

interface LoginRequest {
  email: string;
  password: string;
  name: string;
}

export function SignupForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const router = useRouter();
  const formik = useFormik({
    initialValues: {
      name: 'Anita Williams',
      email: 'anitawilliams@example.com',
      password: 'Anitawilliams123'
    },
    onSubmit: values => handleRegister(values)
  })

  const handleRegister = async (values: LoginRequest) => {
    const resp = await fetch('/api/auth/register', {
      method: 'POST',
      body: JSON.stringify(values)
    });
    const res = await resp.json();
    if (resp.status === 200) {
      return router.push('/login')
    }
  };


  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader className="text-center">
          <CardTitle className="text-xl">Create your account</CardTitle>
          <CardDescription>
            Enter your email below to create your account
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={formik.handleSubmit}>
            <FieldGroup>
              <Field>
                <FieldLabel htmlFor="name">Full Name</FieldLabel>
                <Input
                    id="name"
                    type="text"
                    placeholder="John Doe"
                    required
                    name="name"
                    onChange={formik.handleChange}
                    value={formik.values.name}/>
              </Field>
              <Field>
                <FieldLabel htmlFor="email">Email</FieldLabel>
                <Input
                  id="email"
                  type="email"
                  placeholder="m@example.com"
                  required
                  name="email"
                  onChange={formik.handleChange}
                  value={formik.values.email}
                />
              </Field>
              <Field>
                <Field>
                  <FieldLabel htmlFor="password">Password</FieldLabel>
                  <Input
                      id="password"
                      type="password"
                      required
                      name="password"
                      onChange={formik.handleChange}
                      value={formik.values.password}/>
                </Field>
                <FieldDescription>
                  Must be at least 8 characters long.
                </FieldDescription>
              </Field>
              <Field>
                <Button type="submit">Create Account</Button>
                <FieldDescription className="text-center">
                  Already have an account? <a href="/login">Sign in</a>
                </FieldDescription>
              </Field>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
