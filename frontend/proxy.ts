import {NextRequest, NextResponse} from "next/server";

export function proxy(req: NextRequest) {
    const session = req.cookies.get("session_id");
    if (!session) {
        const loginUrl = new URL("/login", req.url);
        return NextResponse.redirect(loginUrl);
    }

    return NextResponse.next();
}

export const config = {
    matcher: ["/dashboard"],
};
