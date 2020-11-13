// Setup mock object for local storage.
const localStorageMock = (() => {
  let store = {};
  return {
    getItem: (key) => store[key],
    setItem: (key, value) => {
      store[key] = value.toString();
    },
    clear: () => {
      store = {};
    },
    removeItem: (key) => {
      delete store[key];
    },
  };
})();

const analyticsMock = (() => ({
  page: () => { },
}))();

const mockPixieFlags = {
  AUTH0_DOMAIN: '', AUTH0_CLIENT_ID: '', DOMAIN_NAME: '', SEGMENT_UI_WRITE_KEY: '', LD_CLIENT_ID: '',
};

Object.defineProperty(window, 'localStorage', { value: localStorageMock });
Object.defineProperty(window, 'analytics', { value: analyticsMock });
Object.defineProperty(window, '__PIXIE_FLAGS__', { value: mockPixieFlags });

// This is a hack to get clsx to actually work in the test.
jest.mock('clsx', () => ({
  __esModule: true,
  default: jest.requireActual('clsx')
}));

// This prevents console errors about the use of useLayoutEffect on the server
jest.mock('react', () => ({
  ...jest.requireActual('react'),
  useLayoutEffect: jest.requireActual('react').useEffect,
}));