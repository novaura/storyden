import { PropsWithChildren } from "react";

import { OAuthProviderList } from "@/components/auth/OAuthProviderList";
import { LinkButton } from "@/components/ui/link-button";
import { HStack, VStack } from "@/styled-system/jsx";

export default async function Layout({ children }: PropsWithChildren) {
  return (
    <VStack w="full">
      {children}

      <OAuthProviderList />

      <HStack>
        <LinkButton size="xs" variant="ghost" href="/password-reset">
          Forgot password
        </LinkButton>

        <LinkButton size="xs" variant="subtle" href="/register">
          Register
        </LinkButton>
      </HStack>
    </VStack>
  );
}
