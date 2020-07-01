import Autocomplete from 'components/autocomplete/autocomplete';
import { CompletionHeader, CompletionItem } from 'components/autocomplete/completions';
import PixieCommandIcon from 'components/icons/pixie-command';
import { ScriptsContext } from 'containers/App/scripts-context';
import * as React from 'react';
import { argsForVis } from 'utils/args-utils';

import { createStyles, makeStyles } from '@material-ui/core';
import Card from '@material-ui/core/Card';
import Modal from '@material-ui/core/Modal';

import { ExecuteContext } from '../../context/execute-context';
import { LiveViewPage } from '../../components/live-widgets/utils/live-view-params';
import { parseVis } from '../live/vis';

interface CommandInputProps {
  open: boolean;
  onClose: () => void;
}

const useStyles = makeStyles(() => createStyles({
  card: {
    position: 'absolute',
    width: '760px',
    top: '40%',
    left: '50%',
    transform: 'translate(-50%, -20vh)',
  },
  input: {
    maxHeight: '60vh',
  },
}));

// TODO(malthus): Figure out the lifecycle of this component. When a command is selected,
// should the component clear the input? What about when the input is dismised?

const CommandInput: React.FC<CommandInputProps> = ({ open, onClose }) => {
  const classes = useStyles();

  const { scripts } = React.useContext(ScriptsContext);
  const [completions, setCompletions] = React.useState<CompletionItem[]>([]);

  React.useEffect(() => {
    const visibleScripts = [...scripts.entries()]
      .filter(([, s]) => !s.hidden)
      .map(([, s]) => s);

    setCompletions(visibleScripts.map((script) => ({
      type: 'item',
      id: script.id,
      title: script.id,
      description: script.description,
    })));
  }, [scripts]);

  const { execute } = React.useContext(ExecuteContext);

  const getCompletions = React.useCallback((input) => {
    if (!input) {
      return Promise.resolve([{ type: 'header', header: 'Example Scripts' } as CompletionHeader, ...completions]);
    }
    return Promise.resolve(completions.filter((completion) => completion.title.includes(input)));
  }, [completions]);

  const selectScript = (id) => {
    const script = scripts.get(id);
    const vis = parseVis(script.vis);
    if (script) {
      execute({
        liveViewPage: LiveViewPage.Default,
        entityParamNames: [],
        pxl: script.code,
        vis,
        id: script.id,
        // Fill the default args for now. This will go away once the autocomplete is implemented.
        args: argsForVis(vis, {}),
      });
    }
    onClose();
  };

  return (
    <Modal open={open} onClose={onClose} BackdropProps={{}}>
      <Card className={classes.card}>
        <Autocomplete
          className={classes.input}
          placeholder='Pixie Command'
          prefix={<PixieCommandIcon />}
          onSelection={selectScript}
          getCompletions={getCompletions}
        />
      </Card>
    </Modal>
  );
};

export default CommandInput;