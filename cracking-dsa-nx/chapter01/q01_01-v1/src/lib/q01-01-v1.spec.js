import {
  isUnique_OldForNoDataStructures,
  isUnique_OldForWithSet,
  isUnique_ForOfWithSet,
} from './q01-01-v1.js';

describe('Chapter 01, Interview Question 01 -- isUnique', () => {
  describe('old for-loop style', () => {
    describe('with set', () => {
      it('reports the correct result for simple cases', () => {
        expect(isUnique_OldForWithSet('abc')).toBe(true);
        expect(isUnique_OldForWithSet('abca')).toBe(false);
      })

      it('gives the wrong answer with unicode code points', () => {
          // code points share the same first code unit
          expect(() => {
            expect(isUnique_OldForWithSet('🙉💩')).toBe(true);
          }, 'Expected an exception from the test assertion, but none was thrown').toThrowError();
      })
    })

    describe('no additional data structures', () => {
      it('reports the correct result for simple cases', () => {
        expect(isUnique_OldForNoDataStructures('abc')).toBe(true);
        expect(isUnique_OldForNoDataStructures('abca')).toBe(false);
      })

      it('gives the wrong answer with unicode code points', () => {
          // code points share the same first code unit
          expect(() => {
            expect(isUnique_OldForNoDataStructures('🙉💩')).toBe(true);
          }, 'Expected an exception from the test assertion, but none was thrown').toThrowError();
      })
    });
  });

  describe('for-of approach', () => {
    describe('with set', () => {
      it('works for simple cases (code point fits into a single code unit)', () => {
        expect(isUnique_ForOfWithSet('abc')).toBe(true);
        expect(isUnique_ForOfWithSet('abca')).toBe(false);
      })

      it('works with multi-code-unit code points', () => {
        // code points share the same first code unit
        expect(isUnique_ForOfWithSet('🙉💩')).toBe(true);
      })
    })
  });
})
