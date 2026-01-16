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
import * as Yup from "yup";

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
  const validationSchema = Yup.object({
    name: Yup.string()
      .required("Full name is required"),
    email: Yup.string()
      .email("Invalid email address")
      .required("Email is required"),
    password: Yup.string()
      .min(8, "Password must be at least 8 characters")
      .required("Password is required"),
  });
  const formik = useFormik({
    initialValues: {
      name: '',
      email: '',
      password: ''
    },
    validationSchema,
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
                    name="name"
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    value={formik.values.name}/>
                {formik.touched.name && formik.errors.name && (
                  <FieldDescription className="text-red-500 text-sm mt-1">
                    {formik.errors.name}
                  </FieldDescription>
                )}
              </Field>
              <Field>
                <FieldLabel htmlFor="email">Email</FieldLabel>
                <Input
                  id="email"
                  type="email"
                  placeholder="m@example.com"
                  name="email"
                  onChange={formik.handleChange}
                  onBlur={formik.handleBlur}
                  value={formik.values.email}
                />
                {formik.touched.email && formik.errors.email && (
                  <FieldDescription className="text-red-500 text-sm mt-1">
                    {formik.errors.email}
                  </FieldDescription>
                )}
              </Field>
              <Field>
                <Field>
                  <FieldLabel htmlFor="password">Password</FieldLabel>
                  <Input
                      id="password"
                      type="password"
                      name="password"
                      onChange={formik.handleChange}
                      onBlur={formik.handleBlur}
                      value={formik.values.password}/>
                </Field>
                <FieldDescription>
                  Must be at least 8 characters long.
                </FieldDescription>
                {formik.touched.password && formik.errors.password && (
                  <FieldDescription className="text-red-500 text-sm mt-1">
                    {formik.errors.password}
                  </FieldDescription>
                )}
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
