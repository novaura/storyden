import { Account } from "@/api/openapi-schema";
import { nodeList } from "@/api/openapi-server/nodes";
import { threadList } from "@/api/openapi-server/threads";
import { getServerSession } from "@/auth/server-session";
import { FeedConfig } from "@/components/feed/FeedConfig/FeedConfig";
import { UnreadyBanner } from "@/components/site/Unready";
import { Button } from "@/components/ui/button";
import { Settings } from "@/lib/settings/settings";
import { HStack, LStack } from "@/styled-system/jsx";

import { LibraryFeedScreen } from "./LibraryFeedScreen";
import { ThreadFeedScreen } from "./ThreadFeedScreen";

export type Props = {
  initialSession?: Account;
  initialSettings: Settings;
};

export function FeedScreen({ initialSession, initialSettings }: Props) {
  return (
    <LStack>
      <FeedConfig
        initialSession={initialSession}
        initialSettings={initialSettings}
      />
      <FeedScreenContent initialSettings={initialSettings} />
    </LStack>
  );
}

async function FeedScreenContent({ initialSettings }: Props) {
  const feedConfig = initialSettings.metadata.feed;

  switch (feedConfig.source.type) {
    case "threads":
      return <ThreadFeedScreenContent />;

    case "library":
      return <LibraryFeedScreenContent />;
  }
}

async function ThreadFeedScreenContent() {
  try {
    const threads = await threadList();
    return <ThreadFeedScreen initialData={threads.data} />;
  } catch (e) {
    return <UnreadyBanner error={e} />;
  }
}
async function LibraryFeedScreenContent() {
  try {
    const nodes = await nodeList();
    return <LibraryFeedScreen initialData={nodes.data} />;
  } catch (e) {
    return <UnreadyBanner error={e} />;
  }
}
