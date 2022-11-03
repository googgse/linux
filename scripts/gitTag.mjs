// import built-in modules
import * as childProcess from 'node:child_process';
import * as readline from 'node:readline/promises';
import { stdin as input, stdout as output } from 'node:process';

async function main() {
  const rl = readline.createInterface({ input, output });
  // control flow
  for (; ;) {
    const tag = await rl.question('Please input your tag: ');
    // control flow
    if (tag.startsWith('v')) {
      break
    }
  }
  const confirm = await rl.question(`If you wanna tag v${tag}? (Y/n)`);
  // control flow
  if (confirm === 'Y' || confirm === 'y') {
    childProcess.exec(`git tag -a v${tag} -m 'Release v${tag}'`)
  } else {
    process.exit()
  }
  rl.close();
}

main();
